import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';

import * as authActions from '../../actions/auth.action';

function loginForm(props) {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const onUsernameChange = (e) => {
        const { value } = e.target;
        setUsername(value);
    }

    const onPasswordChange = (e) => {
        const { value } = e.target;
        setPassword(value);
    }

    const onFormSubmit = (e) => {
        e.preventDefault();
        props.loginUser(username, password);
    }

    return (
        <div>
            <form onSubmit={onFormSubmit}>
                Username: <br />
                <input name="username" value={username} onChange={onUsernameChange} /> <br />

                Password: <br />
                <input name="password" value={password} onChange={onPasswordChange} /> <br />

                <button type="submit">Login</button> <br />

                <Link to="/register">Register here</Link>
            </form>
        </div>
    )
}

const mapDispatchToProps = dispatch => ({
    loginUser: (username, password) => dispatch(authActions.loginUser(username, password))
});

export default connect(null, mapDispatchToProps)(loginForm);