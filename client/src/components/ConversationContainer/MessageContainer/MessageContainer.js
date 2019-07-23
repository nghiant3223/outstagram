import React, { Component } from 'react';
import { connect } from 'react-redux';

import { renderMessages } from '../../../utils/dom';
import { genUID } from '../../../utils/lang';

import { Message } from '../../../models/message';
import ContainerContext from './ContainerContext';

import * as roomServices from "../../../services/room.service";

import "./MessageContainer.css";
import Loading from '../../Loading/Loading';

class ChatboxContainer extends Component {
    constructor(props) {
        super(props);
        this.state = {
            isLoading: false,
            messageContent: '',
            messages: props.messages
        }
    }

    componentWillReceiveProps(nextProps) {
        if (this.props.messages !== nextProps.messages) {
            this.setState({ messages: nextProps.messages });
        }
    }

    componentDidUpdate(prevProps) {
        if (this.props.roomIdOrUsername !== prevProps.roomIdOrUsername) {
            this.scrollToBottom();
            this.messageInput.focus();
        }
    }

    componentDidMount() {
        this.scrollToBottom();
        this.messageInput.focus();
    }

    scrollToBottom() {
        this.messageContainer.scrollTop = this.messageContainer.scrollHeight;
    }

    onMessageConainerScroll = async () => {
        const { messages } = this.state;
        const { roomIdOrUsername } = this.props;
        const { scrollTop } = this.messageContainer;

        if (scrollTop == 0) {
            this.setState({ isLoading: true });
            try {
                const { data: { data: { messages: fetchMessages } } } = await roomServices.getMessages(roomIdOrUsername, 20, messages.length);
                if (fetchMessages !== null) {
                    this.setState((prevState) => ({ messages: [...fetchMessages, ...prevState.messages] }));
                    setTimeout(() => {
                        this.messageContainer.scrollTop = 10;
                    }, 500);
                }

            } catch (e) {
                console.log("Cannot fetch more message", e);
            } finally {
                this.setState({ isLoading: false });
            }
        }
    }

    onFormSubmit = (e) => {
        e.preventDefault();

        const { user } = this.props;
        const messageContent = this.messageInput.value;
        const newMessage = new Message(genUID(), user.id, messageContent, true);

        this.setState((prevState) => ({ messages: [...prevState.messages, newMessage] }));
        this.messageInput.value = "";
    }

    // Replace the temporary message by the newly created message
    replaceMessage = (uid, newCreatedMessage) => {
        const { messages } = this.state;
        const message = messages.find(message => message.id === uid);

        if (!message) {
            throw new Error("Message does not exist");
        }

        // Copy property from newCreatedMessage to current message in the state;
        for (var k in newCreatedMessage) {
            // IMPORTANT: Ignore id field to prevent changing Message's key, which cause a new Message is created
            if (k !== "id") {
                message[k] = newCreatedMessage[k];
            }
        }

        this.scrollToBottom();
    }

    render() {
        const { user, roomIdOrUsername } = this.props;
        const { messages, isLoading } = this.state;

        return (
            <ContainerContext.Provider value={{ replaceMessage: this.replaceMessage, roomIdOrUsername: roomIdOrUsername }}>
                <div className="ChatboxContainer">
                    <div className="ChatboxContainer__MessageContainer" onScroll={this.onMessageConainerScroll} ref={el => this.messageContainer = el}>
                        {isLoading && <div className="ChatboxContainer__MessageContainer__Loader"><Loading /></div>}
                        <div style={{ padding: "0.5em" }}>
                            {renderMessages(messages, user.id)}
                        </div>
                    </div>

                    <form className="ChatboxContainer__InputContainer" onSubmit={this.onFormSubmit}>
                        <input className="ChatboxContainer__InputContainer__Input" placeholder="Type message..." ref={el => this.messageInput = el} />
                        <div className="ChatboxContainer__InputContainer__SendBtn">
                            <button>send</button>
                        </div>
                    </form>
                </div>
            </ContainerContext.Provider>
        )
    }
}

const mapStateToProps = ({ authReducer: { user } }) => ({ user });

export default connect(mapStateToProps)(ChatboxContainer);