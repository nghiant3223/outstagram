import React, { Component } from 'react';
import { connect } from 'react-redux';

import { renderMessages } from '../../../utils/dom';
import { genUID } from '../../../utils/lang';

import "./MessageContainer.css";
import { Message } from '../../../models/message';

class ChatboxContainer extends Component {
    constructor(props) {
        super(props);
        this.state = {
            messageContent: '',
            messages: props.messages
        }
    }

    componentWillReceiveProps(nextProps) {
        if (this.props.messages !== nextProps.messages) {
            this.setState({ messages: nextProps.messages });
        }
    }

    componentDidUpdate() {
        this.scrollToBottom();
    }

    componentDidMount() {
        this.scrollToBottom();
    }


    scrollToBottom() {
        this.messageContainer.scrollTop = this.messageContainer.scrollHeight;
    }

    onFormSubmit = (e) => {
        e.preventDefault();

        const { user } = this.props;
        const messageContent = this.messageInput.value;
        const newMessage = new Message(genUID(), user.id, messageContent, true);

        this.setState((prevState) => ({ messages: [...prevState.messages, newMessage] }));
        this.messageInput.value = "";
    }

    render() {
        const { user } = this.props;
        const { messages } = this.state;

        return (
            <div className="ChatboxContainer">
                <div className="ChatboxContainer__MessageContainer" ref={el => this.messageContainer = el}>
                    <div style={{ padding: "0.5em" }}>
                        {renderMessages(messages, user.id)}
                    </div>
                </div>

                <form className="ChatboxContainer__InputContainer" onSubmit={this.onFormSubmit}>
                    <input className="ChatboxContainer__InputContainer__Input" placeholder="Type message..." ref={el => this.messageInput = el} />
                    <div className="ChatboxContainer__InputContainer__SendBtn">
                        <button>SEND</button>
                    </div>
                </form>
            </div>
        )
    }
}

const mapStateToProps = ({ authReducer: { user } }) => ({ user });

export default connect(mapStateToProps)(ChatboxContainer);