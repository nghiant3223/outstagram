import React, { Component } from 'react'
import { Search } from 'semantic-ui-react'
import _ from 'lodash';

import * as userServices from '../../../services/user.service';
import Avatar from '../../Avatar/Avatar';

import "./ContactSearch.css";

const resultRenderer = ({ title }) => {
    const [id, fullname] = title.split(' ');

    if (title === "__feching_data__") {
        return <div className="FechingData">Fetching data...</div>;
    }

    if (title === "__no_results__") {
        return <div className="FechingData">No results</div>;
    }

    return <div className="ResultContainer">
        <div><Avatar userID={id} /></div>
        <div className="Fullname">{fullname}</div>
    </div>
}


const initialState = { isLoading: false, results: [], value: '' }

class ContactSearch extends Component {
    state = initialState

    handleResultSelect = (e, { result }) => {
        const components = result.title.split(' ');

        if (components.length > 0) {
            const [id] = components;
            const { onContactClick } = this.props;
            onContactClick(id);
        }
    }

    handleSearchChange = (e, { value }) => {
        if (value == "") {
            this.setState({ value });
            return;
        }

        this.setState({ isLoading: true, value, results: [{ title: "__feching_data__" }] })
        userServices.searchUser(value).then(({ data: { data } }) => {
            if (data) {
                this.setState({ results: data.map((result) => ({ title: result.id + " " + result.fullname })) })
            } else {
                this.setState({ results: [{ title: "__no_results__" }] });
            }
        }).catch((e) => {
            console.log("Cannot search user");
        }).finally(() => {
            this.setState({ isLoading: false });
        });
    }

    render() {
        const { isLoading, value, results } = this.state

        return (
            <Search
                input={{ icon: 'search', iconPosition: 'left' }}
                loading={isLoading}
                onResultSelect={this.handleResultSelect}
                onSearchChange={_.debounce(this.handleSearchChange, 500, { leading: true })}
                results={results}
                value={value}
                resultRenderer={resultRenderer}
            />
        )
    }
}

export default ContactSearch;