import React, { useState, Component } from 'react';
import ContentEditable from 'react-contenteditable';

import "./PostInput.css";
import Avatar from '../Avatar/Avatar';
import { genUID } from '../../utils/lang';

class PostInput extends Component {
    state = {
        content: ""
    }

    id = genUID();

    componentDidMount() {
        const { focusOnMount } = this.props;

        if (focusOnMount) {
            this.focus();
        }
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
        this.focus();
    }

    onKeyDown = (e) => {
        if (e.key === "Enter") {
            const { onSubmit } = this.props;
            onSubmit(this.state.content);
            this.setState({ content: "" });
        }
    }

    focus = () => {
        document.getElementsByClassName(this.id)[0].focus();
    }

    scrollTo = () => {
        window.scrollTo({ top: this.input.offsetTop, behavior: "smooth" });
    }

    render() {
        const { content } = this.state;
        const { inverted, isCommentInput, placeholder, style: propsStyle, userID } = this.props;

        const style = {
            position: "relative",
            backgroundColor: isCommentInput ? "#f2f3f5" : "unset",
            borderTop: isCommentInput ? "1.25px solid rgba(0, 0, 0, 0.1)" : 0,
            borderBottom: isCommentInput ? "1.25px solid rgba(0, 0, 0, 0.1)" : 0,
        }

        return (
            <div
                className="PostInput"
                ref={el => this.input = el}
                style={propsStyle !== undefined ? { ...style, ...propsStyle } : { ...style }}>
                <div>
                    <Avatar width="2em" userID={userID} />
                </div>
                <div className="PostInput__InputContainer">
                    {content === "" && <div className="PostInput__Placeholder" onClick={this.onPlaceholderClick}>{placeholder}</div>}
                    <ContentEditable className={`PostInput__Input ${this.id}`} onChange={this.onInputChange} onKeyDown={this.onKeyDown} html={content} style={{ backgroundColor: !inverted ? "white" : "rgba(0, 0, 0, 0.075)" }} />
                </div>
            </div >
        )
    }

}

export default PostInput;