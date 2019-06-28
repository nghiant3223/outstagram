import React from 'react';
import { connect } from 'react-redux';
import { Dropdown } from 'semantic-ui-react';

import * as authActions from '../../actions/auth.action';

import './Header.css';
import defaultAvatar from '../../images/x.png';

const Header = (props) => {
    const { user } = props;
    return (
        <header>
            <div className="Header">
                <div className="Header__Left">
                    Outstagram
                </div>

                <div className="Header__Right">

                    <div className="Header__Right__Info">
                        <div className="Header__Right__Info__Avatar" >
                            <img src={defaultAvatar} alt="avatar" />
                        </div>
                        <div>{user.fullname}</div>
                    </div>
                    <Dropdown direction='left'>
                        <Dropdown.Menu>
                            <Dropdown.Item onClick={props.logoutUser} text="Logout" />
                        </Dropdown.Menu>
                    </Dropdown>
                </div>
            </div>
        </header>
    );
}
const mapStateToProps = ({ authReducer: { user } }) => ({ user });

const mapDispatchToProps = (dispatch) => ({
    logoutUser: () => dispatch(authActions.logoutUser())
});

export default connect(mapStateToProps, mapDispatchToProps)(Header);