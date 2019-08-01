import React, { Component } from 'react';
import { connect } from "react-redux";

import Container from "../../components/Container/Container";
import UserInfo from './UserInfo/UserInfo';

import * as userServices from '../../services/user.service';
import * as postServices from '../../services/post.service';

import "./ProfilePage.css";
import ProfileImage from './ProfileImage/ProfileImage';
import CoverImage from './CoverImage/CoverImage';
import Post from '../../components/Post/Post';
import ProfilePagePlaceholder from './Placeholder';

class ProfilePage extends Component {
    state = {
        user: undefined,
        posts: [],
        isLoading: true
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
            this.setState({ isLoading: true });
            const { data: { data: { user } } } = await userServices.getUser(username);
            const { data: { data: { posts } } } = await postServices.getPosts(user.id, 100, 0);
            this.setState({ user, posts: posts || [] });
        } catch (e) {
            this.setState({ user: null });
        } finally {
            this.setState({ isLoading: false });
        }
    }

    render() {
        const { user, posts, isLoading } = this.state;

        // If component is loading
        if (user === undefined || isLoading) {
            return <ProfilePagePlaceholder />
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
            <div>
                <Container className="ProfileSummaryContainer">
                    <div className="ImagesContainer">
                        <CoverImage coverImageID={user.coverImageID} />
                        <ProfileImage userID={user.id} />
                    </div>

                    <UserInfo user={user} />
                </Container>

                <Container className="ProfileBodyContainer" white={false} >
                    <div className="ProfileBodyContainer__PostContainer">
                        {posts.map((post) => <div className="ProfileBodyContainer__PostContainer__Post"><Post {...post} key={post.id} showImageGrid={true} /></div>)}
                    </div>

                    <div className="ProfileBodyContainer__BiographyContainer">

                    </div>
                </Container>
            </div>
        )
    }
}

const mapStateToProps = ({ authReducer: { user } }) => ({ user });

export default connect(mapStateToProps)(ProfilePage);