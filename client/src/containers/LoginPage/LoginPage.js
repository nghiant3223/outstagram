import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Redirect } from 'react-router-dom';

import * as authActions from '../../actions/auth.action';

import './LoginPage.css';

class LoginPage extends Component {
    state = {
        username: '',
        password: ''
    }

    onFieldChanged = (e) => {
        this.setState({ [e.target.name]: e.target.value });
    }

    onFormSubmit = (e) => {
        e.preventDefault();
        this.props.loginUser(this.state.username, this.state.password);
    }

    render() {
        if (this.props.isAuthenticated) {
            return <Redirect to="/" />
        }

        return (
            <div>
                <form onSubmit={this.onFormSubmit}>
                    <input name="username" value={this.state.username} onChange={this.onFieldChanged}/>
                    <input name="password" value={this.state.password} onChange={this.onFieldChanged}/>
                    <button type="submit">Submit</button>
                </form>
            </div>
        );
    }
}

const mapStateToProps = ({ auth: { isAuthenticated } }) => ({ isAuthenticated });

const mapDispatchToProps = dispatch => ({
    loginUser: (username, password) => dispatch(authActions.loginUser(username, password))
});

export default connect(mapStateToProps, mapDispatchToProps)(LoginPage);