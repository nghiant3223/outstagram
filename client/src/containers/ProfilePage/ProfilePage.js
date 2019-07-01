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
        const { userID } = this.props.match.params;
        this.getUser(userID);
    }

    componentDidUpdate(prevProps) {
        const { userID } = this.props.match.params;

        if (userID !== prevProps.match.params.userID) {
            this.getUser(userID);
        }
    }

    async getUser(userID) {
        try {
            const { data: { data: user } } = await userServices.getUser(userID);
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