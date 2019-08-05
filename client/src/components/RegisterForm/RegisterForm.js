import React, { useState } from 'react';
import { Link, Redirect } from 'react-router-dom';
import { Button, Form } from 'semantic-ui-react';

import { registerUser } from '../../services/auth.service';

function registerForm() {
    const [email, setEmail] = useState('');
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [fullname, setFullname] = useState('');
    const [avatar, setAvatar] = useState();
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

    const onAvatarChange = (e) => {
        e.persist();
        setAvatar(e.target.files[0]);
    }

    const formValidation = () => {
        if (avatar === undefined) {
            alert("Avatar mustn't be null");
            return false;
        }

        return true;
    }

    const onFormSubmit = (e) => {
        e.preventDefault();

        if (!formValidation()) {
            return;
        }

        registerUser(fullname, email, username, password, avatar)
            .then(() => {
                setShouldRedirect(true);
            })
            .catch((err) => {
                alert(err.response && err.response.data.message);
            });
    }

    if (shouldRedirect) {
        return <Redirect to="/login" />
    }

    return (

        <div className="FormContainer">
            <div className="FormContainer__Inner">
                <Form onSubmit={onFormSubmit}>
                    <Form.Field>
                        <label>Fullname</label>
                        <input name="fullname" placeholder="Fullname" value={fullname} onChange={onFullnameChange} />
                    </Form.Field>
                    <Form.Field>
                        <label>Email</label>
                        <input name="email" placeholder="Email" value={email} onChange={onEmailChange} />
                    </Form.Field>

                    <Form.Field>
                        <label>Username</label>
                        <input name="username" placeholder="Username" value={username} onChange={onUsernameChange} />
                    </Form.Field>

                    <Form.Field>
                        <label>Password</label>
                        <input name="password" placeholder="Password" type="password" value={password} onChange={onPasswordChange} />
                    </Form.Field>

                    <Form.Field>
                        <label>Avatar</label>
                        <input type="file" multiple onClick={e => e.target.value = null} onChange={onAvatarChange} />
                    </Form.Field>

                    <Button type='submit'>Signup</Button>
                    <Link to="/login">Login here</Link>
                </Form>
            </div>
        </div>
    )
}

export default registerForm;