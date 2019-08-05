import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { Transition } from 'semantic-ui-react';

import Avatar from '../../Avatar/Avatar';
import UserInfoPopup from '../../UserInfoPopup/UserInfoPopup';

import "./Suggestion.css";
import FollowButton from '../../FollowButton/FollowButton';

const DISAPPEAR_DURATION = 500;

class Suggestion extends Component {
    state = { visible: true }

    onClick = () => {
        this.setState({ visible: false });
        setTimeout(() => {
            const { onSuggestionClick, id } = this.props;
            onSuggestionClick(id);
        }, DISAPPEAR_DURATION);
    }

    render() {
        const { fullname, username, id } = this.props;
        const { visible } = this.state;

        return (
            <div>
                <Transition visible={visible} animation='fade' duration={DISAPPEAR_DURATION}>
                    <div className="Suggestion" style={{ display: "flex" }}>
                        <UserInfoPopup username={username} trigger={<Link to={`/${username}`}><Avatar width="2.75rem" userID={id} /></Link>} />
                        <UserInfoPopup username={username} trigger={<div className="PostHeader__Info__Fullname"><Link to={`/${username}`}><div className="Fullname Suggestion__Fullname">{fullname}</div></Link></div>} />
                        <FollowButton userID={id} size="tiny" onClick={this.onClick} />
                    </div>
                </Transition>
            </div>
        )
    }
}

export default Suggestion;