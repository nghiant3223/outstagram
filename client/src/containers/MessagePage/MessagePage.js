import React, { Component } from 'react';

import "./MessagePage.css";

import Container from '../../components/Container/Container';
import ContactContainer from '../../components/ContactContainer/ContactContainer';
import ConversationContainer from '../../components/ConversationContainer/ConversationContainer';

class MessagePage extends Component {
    render() {
        return (
            <Container>
                <div className="MessagePage">
                    <ContactContainer />
                    <ConversationContainer />
                </div>
            </Container>
        )
    }
}

export default MessagePage;