import React from 'react';
import { Placeholder } from 'semantic-ui-react'

import "./Post.css";
import "./PostPlaceholder.css";

export default function PostPlaceholder() {
    return (
        <div className="PostPlaceholder">
            <Placeholder fluid>
                <Placeholder.Header image>
                    <Placeholder.Line length='short' />
                    <Placeholder.Line length='very short' />
                </Placeholder.Header>
            </Placeholder>

            <Placeholder style={{ height: 250, marginTop: 16 }} fluid>
                <Placeholder.Image />
            </Placeholder>

            <Placeholder fluid style={{marginTop: 16 }} >
                <Placeholder.Header>
                <Placeholder.Line length='very short' />
                </Placeholder.Header>

            </Placeholder>
        </div>
    );
}
