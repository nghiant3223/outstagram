import React, { Component } from 'react';
import { Placeholder } from 'semantic-ui-react';
import { connect } from "react-redux";

import Container from "../../components/Container/Container";
import UserInfo from './UserInfo/UserInfo';

import * as userServices from '../../services/user.service';
import * as postServices from '../../services/post.service';

import "./ProfilePage.css";
import ProfileImage from './ProfileImage/ProfileImage';
import CoverImage from './CoverImage/CoverImage';
import Post from '../../components/Post/Post';
import PostPlaceholder from '../../components/Post/PostPlaceholder';
import Avatar from '../../components/Avatar/Avatar';
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
            const { data: { data: user } } = await userServices.getUser(username);
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
            return <div>
                <Container className="ProfileSummaryContainer" style={{ padding: "1em" }}>
                    <Placeholder style={{ height: 250, marginTop: 0 }} fluid>
                        <Placeholder.Image />
                    </Placeholder>

                    <div className="ImagesContainer">
                        <div className="ImagesContainer__Avatar">
                            <Avatar size="big" width="125px" />
                        </div>
                    </div>

                    <div className="InfoContainer">
                        <div className="InfoHeader">
                            <div className="InfoHeader__Fullname" style={{ width: 150 }}>
                                <Placeholder>
                                    <Placeholder.Header>
                                        <Placeholder.Line length="full" className="medium-height" />
                                    </Placeholder.Header>
                                </Placeholder>
                            </div>
                        </div>

                        <div className="InfoItemContainer">
                            {Array(3).fill(0).map((_, index) => <div key={index} className="InfoItem" style={{ width: 50 }}>
                                <Placeholder>
                                    <Placeholder.Header>
                                        <Placeholder.Line length='full' className="medium-height" />
                                    </Placeholder.Header>
                                </Placeholder>
                            </div>)}
                        </div>
                    </div>
                </Container>
                <Container className="ProfileBodyContainer" white={false}>
                    {Array(3).fill(0).map((_, index) => <PostPlaceholder key={index} />)}
                </Container>
            </div>
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
                        <CoverImage />
                        <ProfileImage userID={user.id} />
                    </div>

                    <UserInfo user={user} />
                </Container>

                <Container className="ProfileBodyContainer" white={false}>
                    {posts.map((post) => <Post {...post} key={post.id} showImageGrid={true} />)}
                </Container>
            </div>
        )
    }
}

const mapStateToProps = ({ authReducer: { user } }) => ({ user });

export default connect(mapStateToProps)(ProfilePage);