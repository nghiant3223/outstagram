import React from 'react';
import { Placeholder } from 'semantic-ui-react';
import Avatar from '../../components/Avatar/Avatar';
import Container from '../../components/Container/Container';
import PostPlaceholder from '../../components/Post/PostPlaceholder';

export default function ProfilePagePlaceholder() {
    return (
        <div>
            <Container className="ProfileSummaryContainer" style={{ padding: "1em" }}>
                <Placeholder style={{ height: 250, marginTop: 0 }} fluid>
                    <Placeholder.Image />
                </Placeholder>

                <div className="ImagesContainer">
                    <div className="ImagesContainer__Avatar">
                        <Avatar size="big" width="125px" />
                    </div>
                </div>

                <div className="InfoContainer">
                    <div className="InfoHeader">
                        <div className="InfoHeader__Fullname" style={{ width: 150 }}>
                            <Placeholder>
                                <Placeholder.Header>
                                    <Placeholder.Line length="full" className="medium-height" />
                                </Placeholder.Header>
                            </Placeholder>
                        </div>
                    </div>

                    <div className="InfoItemContainer">
                        {Array(3).fill(0).map((_, index) => <div key={index} className="InfoItem" style={{ width: 50 }}>
                            <Placeholder>
                                <Placeholder.Header>
                                    <Placeholder.Line length='full' className="medium-height" />
                                </Placeholder.Header>
                            </Placeholder>
                        </div>)}
                    </div>
                </div>
            </Container>

            <Container className="ProfileBodyContainer" white={false}>
                {Array(3).fill(0).map((_, index) => <PostPlaceholder key={index} />)}
            </Container>
        </div>
    );
}
