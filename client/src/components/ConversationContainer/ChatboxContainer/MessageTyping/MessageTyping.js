import React, { Component, Fragment } from 'react';

import { Popup } from 'semantic-ui-react';
import Avatar from "../../../Avatar/Avatar";
import Typing from '../../../Typing/Typing';
import Socket from '../../../../Socket';

class MessageTyping extends Component {
    state = {
        typingIDs: [] // Should be 
    }

    componentDidMount() {
        Socket.on("ROOM.SERVER.TYPING", this.onSomeoneTyping);
        Socket.on("ROOM.SERVER.STOP_TYPING", this.onSomeoneStopTyping);
    }

    componentWillUnmount() {
        Socket.removeListener("ROOM.SERVER.TYPING", this.onSomeoneTyping);
        Socket.removeListener("ROOM.SERVER.STOP_TYPING", this.onSomeoneStopTyping);
    }

    onSomeoneTyping = (message) => {
        const { roomID } = this.props;
        const { actorID, data } = message;
        if (data.targetRoomID === roomID) {
            this.setState((prevState) => ({ typingIDs: [...prevState.typingIDs, actorID] }));
        }
    }

    onSomeoneStopTyping = (message) => {
        const { roomID } = this.props;
        const { actorID, data } = message;
        if (data.targetRoomID === roomID) {
            this.setState((prevState) => ({ typingIDs: prevState.typingIDs.filter((id) => id !== actorID) }));
        }
    }

    render() {
        const { typingIDs } = this.state;

        return (
            <Fragment>
                {typingIDs.map((id) =>
                    <div key={id} className={["MessageGroup", "MessageGroup--Left"].join(" ")}>
                        <div className="MessageGroup__Avatar"><Avatar userID={id} width="2.25em" /> </div>
                        <div className={["MessageGroup__ChatboxContainer", "MessageGroup__ChatboxContainer--Left"].join(" ")}>
                            <div className="Message">
                                <Popup content={`Someone is typing`} size="mini" position="right center" inverted
                                    trigger={<div className="Message__Content"><Typing /></div>} style={{ padding: "0.75em" }} />
                            </div>
                        </div>
                    </div>)}
            </Fragment>
        );
    }
}

export default MessageTyping;