import React, { Component } from 'react';
import { withRouter } from 'react-router';

import './Search.css';
import { withRouterInnerRef } from '../../../hocs/withRouter';

class Search extends Component {
    state = {
        searchValue: ""
    }

    componentDidUpdate(prevProps) {
        if (!window.location.pathname.includes("/search")) {
            this.input.value = "";
        }

    }

    onSearchInputChange = (e) => {
        this.setState({ searchValue: e.target.value });
    }

    getSearchValue = () => {
        this.input.blur();
        return this.state.searchValue;
    }

    render() {
        const { searchValue } = this.state;

        return (
            <div className="ui small action input">
                <input value={searchValue} ref={el => this.input = el} onChange={this.onSearchInputChange} placeholder="Search..." type="text" />
                <button className="ui icon button">
                    <i aria-hidden="true" className="search icon"></i>
                </button>
            </div>
        );
    }
}

export default withRouterInnerRef(Search);