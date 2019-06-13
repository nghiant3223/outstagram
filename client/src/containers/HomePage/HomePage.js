import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';

import { logoutUser } from '../../actions/auth.action';

import './HomePage.css';

class HomePage extends Component {
    render() {
        const { username, password } = this.props.user;
        return (
            <div>
                Homepage
                {username}: {password}
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
