import React from 'react';
import { connect } from 'react-redux';
import { Link, NavLink } from 'react-router-dom';
import { Dropdown, Popup } from 'semantic-ui-react';

import Container from '../Container/Container';
import Avatar from '../Avatar/Avatar';

import * as authActions from '../../actions/auth.action';
import * as creatorActions from "../../actions/creator.action";

import './Header.css';

import chatIcon from '../../images/chat.png';
import globalIcon from '../../images/globe.png';
import cameraIcon from '../../images/camera.png';

const Header = (props) => {
    const { user, openCreatorModal, logoutUser } = props;

    return (
        <div className="HeaderContainer">
            <Container className="Header" center={false} white={false}>
                <Link to="/">
                    <div className="Header__Left">
                        Outstagram
                     </div>
                </Link>

                <ul className="Header__Right">
                    <li className="Header__Right__Item Header__IconContainer">
                        <Popup content="Upload your images" size="small" inverted trigger={
                            <NavLink to="/" onClick={openCreatorModal}>
                                <img src={cameraIcon} width="20" />
                            </NavLink>} />

                        <Popup content="Messages" size="small" inverted trigger={
                            <NavLink to="/messages" >
                                <img src={chatIcon} width="20" />
                            </NavLink>} />

                        <Popup content="Notifications" size="small" inverted trigger={
                            <NavLink>
                                <img src={globalIcon} width="20" />
                            </NavLink>} />

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

                        <Dropdown direction='left'>
                            <Dropdown.Menu>
                                <Dropdown.Item onClick={logoutUser} text="Logout" />
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
    logoutUser: () => dispatch(authActions.logoutUser()),
    openCreatorModal: () => dispatch(creatorActions.openModal())
});

export default connect(mapStateToProps, mapDispatchToProps)(Header);