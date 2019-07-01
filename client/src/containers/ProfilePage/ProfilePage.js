import React, { Component } from 'react';

import "./ProfilePage.css";
import Container from "../../components/Container/Container";
import UserInfo from './UserInfo/UserInfo';

class ProfilePage extends Component {
    render() {
        return (
            <Container>
                <div className="ImagesContainer">
                    <div className="ImagesContainer__Cover"></div>
                    <div className="ImagesContainer__Avatar"></div>
                </div>

                <UserInfo />
            </Container>
        )
    }
}

export default ProfilePage;