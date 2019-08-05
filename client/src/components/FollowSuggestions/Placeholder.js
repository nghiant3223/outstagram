import React from 'react';
import { Placeholder } from 'semantic-ui-react';

export default function FollowSuggestionPlaceholder() {
    return (
        <div>
            <Placeholder style={{ height: "1.5em", width: 200 }}>
                <Placeholder.Header>
                </Placeholder.Header>
            </Placeholder>

            <Placeholder fluid>
                <Placeholder.Header image>
                    <Placeholder.Line length='long' />
                    <Placeholder.Line length='medium' />
                </Placeholder.Header>
            </Placeholder>

            <Placeholder fluid>
                <Placeholder.Header image>
                    <Placeholder.Line length='long' />
                    <Placeholder.Line length='medium' />
                </Placeholder.Header>
            </Placeholder>
        </div>
    );
}
