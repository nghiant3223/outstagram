import React, { useState, Component } from 'react';
import ContentEditable from 'react-contenteditable';

import "./CommentInput.css";
import Avatar from '../Avatar/Avatar';

class CommentInput extends Component {
    state = {
        content: ""
    }

    onInputChange = (e) => {
        const text = e.target.value;

        if (text === "<div><br></div><div><br></div>") {
            this.setState({ content: "" });
            return;
        }

        this.setState({ content: text });
    }

    onPlaceholderClick = () => {
        document.getElementsByClassName("CommentInput__Input")[0].focus();
    }

    onKeyDown = (e) => {
        if (e.key === "Enter") {
            this.setState({ content: "" });
        }
    }

    render() {
        const { content } = this.state;
        const { style = {} } = this.props;

        return (
            <div className="CommentInput" style={{ position: "relative", ...style }}>
                <Avatar width="2em" />
                {content === "" && <div className="CommentInput__Placeholder" onClick={this.onPlaceholderClick}>Write your comment...</div>}
                <ContentEditable className="CommentInput__Input" onChange={this.onInputChange} onKeyDown={this.onKeyDown} html={content} />
            </div>
        )
    }

}

export default CommentInput;