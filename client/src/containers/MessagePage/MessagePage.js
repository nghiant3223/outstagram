import React, { Component } from 'react';

import "./MessagePage.css";
import Container from '../../components/Container/Container';
import ContactContainer from '../../components/ContactContainer/ContactContainer';
import ChatboxContainer from '../../components/ChatboxContainer/ChatboxContainer';
import ContactInfo from '../../components/ChatboxContainer/ContactInfo/ContactInfo';

class MessagePage extends Component {
    render() {
        return (
            <Container >
                <div className="MessagePage">
                    <ContactContainer />
                    <div className="MessageInfoContainer">
                        <ContactInfo />
                        <ChatboxContainer />
                    </div>
                </div>
            </Container>
        )
    }
}

export default MessagePage;