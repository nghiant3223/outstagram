import React, { Component } from 'react';

import MessageGroup from './MessageGroup/MessageGroup';

import "./ChatboxContainer.css";
import sendIcon from "../../images/right-arrow.png";

class ChatboxContainer extends Component {
    render() {
        return (
            <div className="ChatboxContainer">
                <div className="ChatboxContainer__MessageContainer">
                    <div style={{ padding: "0 0.5em" }}>
                        <MessageGroup messages={[{
                            content: '23213',
                            createdAt: new Date()
                        }, {
                            content: 'lmaooooooooooooooooooooooooooooooooooooooooooooooo',
                            createdAt: new Date()
                        }]} />

                        <MessageGroup right messages={[{
                            content: '23213',
                            createdAt: new Date()
                        }, {
                            content: 'lmaooooooooooooooooooooooooooooooooooooooooooooooo',
                            createdAt: new Date()
                        }]} />
                        <MessageGroup right messages={[{
                            content: '23213',
                            createdAt: new Date()
                        }, {
                            content: 'lmaooooooooooooooooooooooooooooooooooooooooooooooo',
                            createdAt: new Date()
                        }]} />                    <MessageGroup right messages={[{
                            content: '23213',
                            createdAt: new Date()
                        }, {
                            content: 'lmaooooooooooooooooooooooooooooooooooooooooooooooo',
                            createdAt: new Date()
                        }]} />
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