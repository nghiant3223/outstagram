import React, { Component } from 'react';
import { Redirect } from "react-router-dom";

import Container from "../../components/Container/Container";
import UserInfo from './UserInfo/UserInfo';

import * as userServices from '../../services/user.service';

import "./ProfilePage.css";
import ProfileImage from './ProfileImage/ProfileImage';
import CoverImage from './CoverImage/CoverImage';

class ProfilePage extends Component {
    state = {
        user: undefined
    }

    componentDidMount() {
        const { username } = this.props.match.params;
        this.getUser(username);
    }

    componentDidUpdate(prevProps) {
        const { username } = this.props.match.params;

        if (username !== prevProps.match.params.username) {
            this.getUser(username);
        }
    }

    async getUser(username) {
        try {
            const { data: { data: user } } = await userServices.getUser(username);
            this.setState({ user })
        } catch (e) {
            this.setState({ user: null });
        }
    }

    render() {
        const { user } = this.state;

        // If component is loading
        if (user === undefined) {
            return null;
        }

        // If user is not found
        if (user === null) {
            return (
                <Container>
                    <div>User not found</div>
                </Container>
            );
        }

        return (
            <Container>
                <div className="ImagesContainer">
                    <CoverImage />
                    <ProfileImage />
                </div>

                <UserInfo user={this.state.user} />
            </Container>
        )
    }
}

export default ProfilePage;