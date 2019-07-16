import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';

import "./LoginForm.css";

import * as authActions from '../../actions/auth.action';
import { Button, Form } from 'semantic-ui-react';

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
        <div className="FormContainer">
            <div className="FormContainer__Inner">
                <Form onSubmit={onFormSubmit}>
                    <Form.Field>
                        <label>Username</label>
                        <input name="username" placeholder="Username" value={username} onChange={onUsernameChange} /> <br />
                    </Form.Field>
                    <Form.Field>
                        <label>Password</label>
                        <input name="password" placeholder="Password" type="password" value={password} onChange={onPasswordChange} /> <br />
                    </Form.Field>
                    <Button type='submit'>Login</Button>
                    <Link to="/register">Register here</Link>
                </Form>
            </div>
        </div>
    )
}

const mapDispatchToProps = dispatch => ({
    loginUser: (username, password) => dispatch(authActions.loginUser(username, password))
});

export default connect(null, mapDispatchToProps)(loginForm);