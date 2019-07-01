import React, { Component } from 'react';
import { Link } from 'react-router-dom';

import './NotFoundPage.css';
import Container from '../../components/Container/Container';

class NotFoundPage extends Component {
    render() {
        return (
            <Container>
                <div className="text-center">
                    <span className="error" id="error-code">Not found</span>
                    <Link to="/">Back to homepage</Link>
                </div>
            </Container>
        );
    }
}

export default NotFoundPage;