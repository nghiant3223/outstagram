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
class ProfilePage extends Component {
    state = {
        user: undefined,
        posts: []
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
            const { data: { data: { posts } } } = await postServices.getPosts(user.id, 100, 0);
            this.setState({ user, posts: posts || [] });
        } catch (e) {
            this.setState({ user: null });
        }
    }

    render() {
        const { user, posts } = this.state;

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

        const images = ['https://unsplash.it/100/400',
            'https://www.gettyimages.ie/gi-resources/images/Homepage/Hero/UK/CMS_Creative_164657191_Kingfisher.jpg',
            // 'https://images.pexels.com/photos/1133957/pexels-photo-1133957.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500',
            // 'https://cdn.pixabay.com/photo/2016/10/27/22/53/heart-1776746_960_720.jpg',
            // 'https://images.pexels.com/photos/257840/pexels-photo-257840.jpeg?auto=compress&cs=tinysrgb&h=350',
            // "https://images.pexels.com/photos/67636/rose-blue-flower-rose-blooms-67636.jpeg?auto=compress&cs=tinysrgb&h=350",
            "https://cdn.pixabay.com/photo/2015/04/19/08/32/rose-729509__340.jpg"]


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
                    {posts.map((post) => <Post images={images} {...post} key={post.id} />)}
                </Container>
            </div>
        )
    }
}

const mapStateToProps = ({ authReducer: { user } }) => ({ user });

export default connect(mapStateToProps)(ProfilePage);