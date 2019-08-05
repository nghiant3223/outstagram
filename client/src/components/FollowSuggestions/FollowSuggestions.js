import React, { Component } from 'react'
import Suggestion from './Suggestion/Suggestion';

import * as userServices from '../../services/user.service';
import { Header } from 'semantic-ui-react';
import FollowSuggestionPlaceholder from './Placeholder';

class FollowSuggestions extends Component {
    state = {
        isLoading: true,
        suggestions: [],
        onDisplaySuggestions: []
    }

    componentDidMount() {
        this.setState({ isLoading: true });
        userServices.getSuggestion()
            .then(({ data: { data: users } }) => {
                this.setState({
                    onDisplaySuggestions: users.slice(0, 3),
                    suggestions: users.slice(3)
                });
            }).catch((e) => {
                console.log(e);
            }).finally(() => {
                this.setState({ isLoading: false });
            });
    }

    onSuggestionClick = (id) => {
        this.setState((prevState) => {
            const { suggestions, onDisplaySuggestions } = prevState;

            if (suggestions.length === 0) return prevState;

            const newOnDisplaySuggestions = onDisplaySuggestions.filter((s) => s.id !== id);
            newOnDisplaySuggestions.unshift(suggestions[0]);
            const newSuggestions = suggestions.slice(1);
            return { suggestions: newSuggestions, onDisplaySuggestions: newOnDisplaySuggestions };
        });
    }

    render() {
        const { onDisplaySuggestions, isLoading } = this.state;

        if (isLoading) return <FollowSuggestionPlaceholder />

        return (
            <div>
                <Header as="h5">People you may know</Header>
                {onDisplaySuggestions.map((s) => <Suggestion key={s.id} id={s.id} fullname={s.fullname} username={s.username} onSuggestionClick={this.onSuggestionClick} />)}
            </div>
        )
    }
}

export default FollowSuggestions;