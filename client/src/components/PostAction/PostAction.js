import React from 'react';

import "./PostAction.css";
import { Icon, Grid } from 'semantic-ui-react';

function PostAction({ style = {} }) {
    return (
        <div className="PostAction" style={{ ...style }}>
            <Grid columns={2} stackable textAlign='center'>
                <Grid.Row verticalAlign='middle'>
                    <Grid.Column>
                        <div><Icon name="heart outline" color="black" inverted />Like</div>
                    </Grid.Column>

                    <Grid.Column>
                        <div><Icon name="comment outline" color="black" inverted /> Comment</div>
                    </Grid.Column>
                </Grid.Row>
            </Grid>
        </div>
    )
}

export default PostAction;