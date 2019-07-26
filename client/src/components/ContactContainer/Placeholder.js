import React from 'react';
import _ from 'lodash';
import { Placeholder } from 'semantic-ui-react';

export default function ContactContainerPlaceHolder() {
    return (
        <div className="ContactContainer">
            <div className="ContactContainer__SearchContainer">
                <Placeholder style={{ height: 25, width: "50%", marginTop: 0 }}>
                    <Placeholder.Header>
                        <Placeholder.Image />
                    </Placeholder.Header>
                </Placeholder>
            </div>
            <div className="ContactContainer__ContactItemContainer">
                <div style={{ padding: "1em" }}>
                    {_.times(6, String).map(i => (
                        <Placeholder key={i}>
                            <Placeholder.Header image>
                                <Placeholder.Line length='medium' />
                                <Placeholder.Line length='very long' />
                            </Placeholder.Header>
                        </Placeholder>))}
                </div>
            </div>
        </div>
    )
}