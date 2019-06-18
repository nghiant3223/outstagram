import React, { useState } from 'react';
import { Link, Redirect } from 'react-router-dom';

import { registerUser } from '../../services/auth.service';

function registerForm(props) {
    const [email, setEmail] = useState('');
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [fullname, setFullname] = useState('');
    const [shouldRedirect, setShouldRedirect] = useState(false);

    const onFullnameChange = (e) => {
        const { value } = e.target;
        setFullname(value);
    }

    const onUsernameChange = (e) => {
        const { value } = e.target;
        setUsername(value);
    }

    const onPasswordChange = (e) => {
        const { value } = e.target;
        setPassword(value);
    }

    const onEmailChange = (e) => {
        const { value } = e.target;
        setEmail(value);
    }

    const onFormSubmit = (e) => {
        e.preventDefault();
        registerUser(fullname, email, username, password)
            .then((res) => {
                setShouldRedirect(true);
            })
            .catch((err) => {
                alert(err);
            });
    }

    if (shouldRedirect) {
        return <Redirect to="/login" />
    }

    return (

        <div>
            <form onSubmit={onFormSubmit}>
                Fullname: <br />
                <input name="fullname" value={fullname} onChange={onFullnameChange} /> <br />

                Email: <br />
                <input name="email" value={email} onChange={onEmailChange} /> <br />

                Username: <br />
                <input name="username" value={username} onChange={onUsernameChange} /> <br />

                Password: <br />
                <input name="password" value={password} onChange={onPasswordChange} /> <br />

                <button type="submit">Register</button> <br />

                <Link to="/login">Login here</Link>
            </form>
        </div>
    )
}

export default registerForm;