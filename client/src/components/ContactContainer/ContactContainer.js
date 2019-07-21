import React, { Component } from 'react'
import ContactSearch from './Search/ContactSearch';

import "./ContactContainer.css";
import Avatar from '../Avatar/Avatar';

class ContactContainer extends Component {
    render() {
        return (
            <div className="ContactContainer">
                <div className="ContactContainer__SearchContainer">
                    <ContactSearch />
                </div>
                <div className="ContactContainer__ContactItemContainer">
                    <div className="ContactItemContainer__ContactItem">
                        <div>
                            <Avatar />
                        </div>
                        <div className="ContactItemContainer__ContactItem__Content">
                            <div className="Fullname">Trọng Nghĩa</div>
                            <div>lmaoooooooooooooooooooooooooooooooo</div>
                        </div>
                        <div className="ContactItemContainer__ContactItem__Time">
                            1.12
                        </div>
                    </div>

                    <div className="ContactItemContainer__ContactItem ContactItemContainer__ContactItem--Active">
                        <div>
                            <Avatar />
                        </div>
                        <div className="ContactItemContainer__ContactItem__Content">
                            <div className="Fullname">Trọng Nghĩa</div>
                            <div>lmaoooooooooooooooooooooooooooooooo</div>
                        </div>
                        <div className="ContactItemContainer__ContactItem__Time">
                            1.12
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ContactContainer;