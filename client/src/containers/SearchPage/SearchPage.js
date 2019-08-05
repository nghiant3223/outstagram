import React, { Component, Fragment } from 'react';
import { Link } from 'react-router-dom';
import _ from 'lodash';
import moment from 'moment'

import Container from '../../components/Container/Container';
import * as userServices from '../../services/user.service';
import * as postServices from '../../services/post.service';

import './SearchPage.css';
import { Placeholder, List } from 'semantic-ui-react';
import UserInfoPopup from '../../components/UserInfoPopup/UserInfoPopup';
import Avatar from '../../components/Avatar/Avatar';
import FollowButton from '../../components/FollowButton/FollowButton';
import GridImageContainer from '../../components/GridImageContainer/GridImageContainer';
import FeedbackSummary from '../../components/FeedbackSummary/FeedbackSummary';

class SearchPage extends Component {
    state = {
        users: [],
        posts: [],
        isLoading: true
    }

    componentDidMount() {
        this.fetchResult();
    }

    componentDidUpdate(prevProps) {
        if (this.props.location.search !== prevProps.location.search) {
            this.fetchResult();
        }
    }

    fetchResult() {
        const queryString = this.props.location.search;
        const params = new URLSearchParams(queryString);
        const q = params.get("q");

        this.setState({ isLoading: true });
        Promise.all([userServices.searchUser(q), postServices.searchPost(q)])
            .then(([{ data: { data: userResults } }, { data: { data: postResults } }]) => {
                this.setState({ users: userResults || [], posts: postResults || [] });
            }).catch((e) => {
                console.log(e);
            }).finally(() => {
                this.setState({ isLoading: false });
            });
    }

    render() {
        const { isLoading } = this.state;

        if (isLoading) return <Container size="standard" white shadow>
            <List celled>
                {_.times(6, String).map(i => (
                    <List.Item key={i}>
                        <div className="SearchPage__Placeholder SearchPage__Card">
                            <Placeholder >
                                <Placeholder.Header image>
                                    <Placeholder.Line length='short' />
                                    <Placeholder.Line length='very short' />
                                </Placeholder.Header>
                            </Placeholder>
                        </div>
                    </List.Item>))}
            </List>
        </Container>

        const { users, posts } = this.state;

        return (
            <Fragment>
                <Container size="standard" white shadow>
                    <List celled>
                        {users.map((user) => (
                            <List.Item key={user.id}>
                                <div className="SearchPage__Result SearchPage__Card">
                                    <div>
                                        <UserInfoPopup username={user.username} trigger={<Link to={`/${user.username}`}><Avatar width="2.75rem" userID={user.id} /></Link>} /></div>
                                    <div>
                                        <UserInfoPopup username={user.username} trigger={<div className="PostHeader__Info__Fullname"><Link to={`/${user.username}`}><div className="Fullname">{user.fullname}</div></Link></div>} /></div>
                                    <div>
                                        {!user.isMe && <FollowButton size="tiny" followed={user.followed} userID={user.id} />}
                                    </div>
                                </div>
                            </List.Item>
                        ))}
                    </List>
                </Container>

                {posts.map((post) => {
                    const { id, imageID, ownerID, ownerFullname, ownerUsername, images, createdAt, content, reactableID, reactors, commentCount, reacted, reactCount } = post;
                    return (
                        <Container key={post.id} size="standard" white shadow style={{ padding: "1em 1em 0.5em 1em", margin: "1em 0" }}>
                            <div className="PostHeader">
                                <div className="PostHeader__Avatar" style={{ marginTop: "0.5em" }}>
                                    <UserInfoPopup username={ownerUsername} trigger={<Link to={`/${ownerUsername}`}><Avatar width="2.75rem" userID={ownerID} /></Link>} />
                                </div>

                                <div className="PostHeader__Info">
                                    <UserInfoPopup username={ownerUsername} trigger={<div className="PostHeader__Info__Fullname"><Link to={`/${ownerUsername}`}><div className="Fullname">{ownerFullname}</div></Link></div>} />
                                    <div className="PostHeader__Info__CreatedAt">
                                        {moment(new Date(createdAt)).calendar()}
                                    </div>
                                </div>
                            </div>

                            <Link to={`/posts/${id}`}>
                                <div className="PostResult">
                                    <div> {content} </div>
                                    <div className="PostResult__ImageContainer">
                                        <GridImageContainer images={images ? images : [{ id, imageID }]} />
                                    </div>
                                </div>
                            </Link>

                            {reactCount + commentCount > 0 && <FeedbackSummary reactableID={reactableID} reactors={reactors} commentCount={commentCount} reacted={reacted} reactCount={reactCount} displayCommentCount={commentCount} />}
                        </Container>
                    )
                })}
            </Fragment>
        )
    }
}

export default SearchPage;