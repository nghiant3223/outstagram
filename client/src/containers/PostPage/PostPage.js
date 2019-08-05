import React, { Component } from 'react';
import { Redirect } from 'react-router'
import _ from 'lodash';

import * as postServices from "../../services/post.service";
import Container from '../../components/Container/Container';
import { Header } from 'semantic-ui-react';
import PostPlaceholder from '../../components/Post/PostPlaceholder';
import Post from '../../components/Post/Post';
import FollowSuggestions from '../../components/FollowSuggestions/FollowSuggestions';

class PostPage extends Component {
    state = {
        post: undefined,
        isLoading: false
    }

    componentDidMount() {
        const { postID } = this.props.match.params;
        this.fetchPost(postID);
    }

    fetchPost(id) {
        this.setState({ isLoading: true });
        postServices.getSpecificPost(id)
            .then(({ data: { data: post } }) => {
                this.setState({ post: post });
            }).catch((e) => {
                this.setState({ post: null });
            }).finally(() => {
                this.setState({ isLoading: false });
            });
    }

    render() {
        return (
            <Container white={false}>
                <div className="HomePage__MainContainer">
                    <Container className="HomePage__PostContainer" white={false} center={false}>
                        {this.renderContent()}
                    </Container>

                    <Container className="HomePage__MainContainer__Aside">
                        <FollowSuggestions />
                    </Container>
                </div>
            </Container>
        )
    }

    renderContent() {
        const { post, isLoading } = this.state;

        if (isLoading) return _.times(1, String).map((i) => <PostPlaceholder key={i} />);

        if (post) return <div className="ProfileBodyContainer__PostContainer__Post"><Post {...post} showImageGrid /></div>

        if (post === null) return <Redirect to="/notfound" />

        return null;
    }
}

export default PostPage;