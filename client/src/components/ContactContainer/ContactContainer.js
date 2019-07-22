import React, { Component } from 'react';
import ContactSearch from './Search/ContactSearch';

import "./ContactContainer.css";
import Contact from './Contact/Contact';
import * as roomServices from "../../services/room.service";

class ContactContainer extends Component {
    state = {
        rooms: [],
        isLoading: false
    }

    async componentDidMount() {
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