import React from 'react'
import { Icon } from 'semantic-ui-react';

import "./FeedbackSummary.css";

function FeedbackSummary({ style = {} }) {
    return (
        <div className="FeedbackSummary" style={{ ...style }}>
            <div className="FeedbackSummary__Left">
                <Icon name="heart" color="red" inverted /> You, Wraith King and 132 others
            </div>

            <div className="FeedbackSummary__Righg">
                13 comments
            </div>
        </div>
    )
}

export default FeedbackSummary;