import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';

import Socket from '../../socket';
import { logoutUser } from '../../actions/auth.action';

import './HomePage.css';

class HomePage extends Component {
    componentDidMount() {
        Socket.get();
    }

    render() {
        const socket = Socket.get();
        const { username, password } = this.props.user;
        return (
            <div>
                Homepage
                <img src="/images/6fa0fe89f10bf85450adc2bd818b9cbd740d9445.png" />
                {username}: {password}
                <button onClick={() => Socket.sendMessage({data:'hello world', type:'STORY.CLIENT.POST_STORY'})}>Send</button>
                <button onClick={this.props.logoutUser}>Logout</button>
            </div>
        );
    }
}

const mapStateToProps = ({ auth: { user } }) => ({ user });

const mapDispatchToProps = (dispatch) => ({
    logoutUser: () => dispatch(logoutUser())
});

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
