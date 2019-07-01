import React from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { Dropdown } from 'semantic-ui-react';

import Container from '../Container/Container';

import * as authActions from '../../actions/auth.action';

import './Header.css';
import defaultAvatar from '../../images/x.png';

const Header = (props) => {
    const { user } = props;
    return (
        <Container className="Header">
            <Link to="/"><div className="Header__Left">
                Outstagram
                </div>
            </Link>

            <div className="Header__Right">
                <Link to={`/${user.id}`}>
                    <div className="Header__Right__Info">
                        <div className="Header__Right__Info__Avatar" >
                            <img src={defaultAvatar} alt="avatar" />
                        </div>
                        <div>{user.fullname}</div>
                    </div>
                </Link>
                <Dropdown direction='left'>
                    <Dropdown.Menu>
                        <Dropdown.Item onClick={props.logoutUser} text="Logout" />
                    </Dropdown.Menu>
                </Dropdown>
            </div>
        </Container>
    );
}
const mapStateToProps = ({ authReducer: { user } }) => ({ user });

const mapDispatchToProps = (dispatch) => ({
    logoutUser: () => dispatch(authActions.logoutUser())
});

export default connect(mapStateToProps, mapDispatchToProps)(Header);