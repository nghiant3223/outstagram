import React from 'react';
import { connect } from 'react-redux';
import { Link, NavLink } from 'react-router-dom';
import { Dropdown, Icon } from 'semantic-ui-react';

import Container from '../Container/Container';

import * as authActions from '../../actions/auth.action';

import './Header.css';
import Avatar from '../Avatar/Avatar';

const Header = (props) => {
    const { user } = props;

    return (
        <div className="HeaderContainer">
            <Container className="Header" center={false} white={false}>
                <Link to="/">
                    <div className="Header__Left">
                        Outstagram
                     </div>
                </Link>

                <ul className="Header__Right">
                    <li className="Header__Right__Item">
                        <NavLink to="/messages" activeStyle={{ color: "blue" }} style={{ color: "grey" }}>
                            <Icon name="mail" size="large" />
                        </NavLink>
                    </li>

                    <li className="Header__Right__Item">
                        <Link to={`/${user.username}`}>
                            <div className="Header__Right__Info">
                                <div className="Header__Right__Info__Avatar" >
                                    <Avatar userID={user.id} width="2em" />
                                </div>
                                <div>{user.fullname}</div>
                            </div>
                        </Link>
                    </li>

                    <li className="Header__Right__Item">
                        <Dropdown direction='left'>
                            <Dropdown.Menu>
                                <Dropdown.Item onClick={props.logoutUser} text="Logout" />
                            </Dropdown.Menu>
                        </Dropdown>
                    </li>
                </ul>
            </Container>
        </div>
    );
}
const mapStateToProps = ({ authReducer: { user } }) => ({ user });

const mapDispatchToProps = (dispatch) => ({
    logoutUser: () => dispatch(authActions.logoutUser())
});

export default connect(mapStateToProps, mapDispatchToProps)(Header);