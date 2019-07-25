import React, { Component } from 'react';
import ContactSearch from './Search/ContactSearch';

import "./ContactContainer.css";
import Contact from './Contact/Contact';
import * as roomServices from "../../services/room.service";
import Socket from '../../Socket';

class ContactContainer extends Component {
    state = {
        rooms: [],
        isLoading: false
    }

    componentDidMount() {
        Socket.on("ROOM.SERVER.SEND_MESSAGE", this.onMessageReceive);
        this.fetchMessages();
    }

    componentWillUnmount() {
        Socket.removeListener("ROOM.SERVER.SEND_MESSAGE", this.onMessageReceive);
    }

    onMessageReceive = (message) => {
        const { targetRoomID } = message.data;

        this.setState((prevState) => {
            let candidateRoom;
            const restRooms = prevState.rooms.filter(room => {
                const match = room.id == targetRoomID;
                if (match) candidateRoom = room;
                return !match;
            });
            return { rooms: [candidateRoom, ...restRooms] };
        });
    }

    async fetchMessages() {
        this.setState({ isLoading: true });
        try {
            const { data: { data: { rooms } } } = await roomServices.getRecentRooms();
            this.setState({ rooms: rooms || [] });
        } catch (e) {
            console.log("Error while fetching rooms", e);
        } finally {
            this.setState({ isLoading: false });
        }
    }

    render() {
        const { rooms } = this.state;

        return (
            <div className="ContactContainer">
                <div className="ContactContainer__SearchContainer">
                    <ContactSearch onContactClick={this.onContactClick} />
                </div>
                <div className="ContactContainer__ContactItemContainer">
                    {rooms.map((contact) => <Contact {...contact} key={contact.id} />)}
                </div>
            </div>
        )
    }
}

export default ContactContainer;