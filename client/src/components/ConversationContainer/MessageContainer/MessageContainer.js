import React, { Component } from 'react';

import MessageGroup from './MessageGroup/MessageGroup';

import "./MessageContainer.css";
import { renderMessages } from '../../../utils/dom';

class ChatboxContainer extends Component {
    componentDidUpdate() {
        this.scrollToBottom();
    }

    componentDidMount() {
        this.scrollToBottom();
    }


    scrollToBottom() {
        this.messageContainer.scrollTop = this.messageContainer.scrollHeight;
        console.log('here', this.messageContainer);
    }

    render() {
        const messages = [{
            id :0,
            authorID: 1,
            content: "123",
            createdAt: 2
        }, {
            id:1,
            authorID: 1,
            content: "123",
            createdAt: 3
        }, {
            id:2,
            authorID: 2,
            content: "123",
            createdAt: 10
        }, {
            id:3,
            authorID: 1,
            content: "123",
            createdAt: 12
        }, {
            id:4,
            authorID: 1,
            content: "123",
            createdAt: 14
        }, {
            id:5,
            authorID: 2,
            content: "123",
            createdAt: 20
        }, {id:6,
            authorID: 1,
            content: "123",
            createdAt: 22
        }, {id:7,
            authorID: 1,
            content: "123",
            createdAt: 12
        }, {id:8,
            authorID: 1,
            content: "123",
            createdAt: 14
        }, {id:9,
            authorID: 2,
            content: "123",
            createdAt: 20
        }, {id:10,
            authorID: 1,
            content: "123",
            createdAt: 22
        }]


        return (
            <div className="ChatboxContainer">
                <div className="ChatboxContainer__MessageContainer"  ref={el => this.messageContainer = el}>
                    <div style={{ padding: "0.5em" }}>
                        {renderMessages(messages, 1)}
                    </div>
                </div>

                <form className="ChatboxContainer__InputContainer">
                    <input className="ChatboxContainer__InputContainer__Input" placeholder="Type message..." />
                    <div className="ChatboxContainer__InputContainer__SendBtn">
                        <button>SEND</button>
                    </div>
                </form>
            </div>
        )
    }
}

export default ChatboxContainer;