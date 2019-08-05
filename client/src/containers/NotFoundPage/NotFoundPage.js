import React, { Component } from 'react';
import { Link } from 'react-router-dom';

import './NotFoundPage.css';
import Container from '../../components/Container/Container';
import { Header, Button } from 'semantic-ui-react';

class NotFoundPage extends Component {
    render() {
        return (
            <Container className="NotFoundContainer">
                <div>404</div>
                <Header as="h3">The page you are looking for is not found</Header>
                <Button color='blue'><Link to="/"><span style={{ color: "white" }}>Go to homepage</span></Link></Button>
            </Container>
        );
    }
}

export default NotFoundPage;