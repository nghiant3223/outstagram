import React, { Component } from 'react';
import { connect } from 'react-redux';

import './HomePage.css';

class HomePage extends Component {
    render() {
        return (
            <div>
                Homepage
            </div>
        );
    }
}

const mapStateToProps = ({ auth: { user } }) => ({ user });

export default connect(mapStateToProps)(HomePage);
