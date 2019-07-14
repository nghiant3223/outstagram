import React, { Component } from 'react';
import { connect } from 'react-redux'
import { Grid } from 'semantic-ui-react';
import PropTypes from 'prop-types';

import * as threaterActions from "../../actions/threater.modal";
import { noAuthStatic } from "../../axios"

import "./GridImageContainer.css";

class GridImageContainer extends Component {
    static defaultProps = {
        images: [],
        hideOverlay: false,
        renderOverlay: () => 'Render text',
        overlayBackgroundColor: '#222222',
        onClickEach: null,
        countFrom: 4
    }

    constructor(props) {
        super(props)

        this.state = {
            modal: false,
            countFrom: props.countFrom > 0 && props.countFrom < 5 ? props.countFrom : 5,
            conditionalRender: false
        };

        this.openModal = this.openModal.bind(this);
        this.onClose = this.onClose.bind(this);

        if (props.countFrom <= 0 || props.countFrom > 5) {
            console.warn('countFrom is limited to 5!')
        }
    }

    openModal(currentIndex) {
        const { openThreaterModal, images } = this.props;
        openThreaterModal({ currentIndex, images });
    }

    onClose() {
        this.setState({ modal: false })
    }

    renderOne() {
        const { countFrom } = this.state;
        const { images } = this.props;
        const imageURLs = images.map(image => noAuthStatic(`/images/others/${image.id}?size=origin`));
        const overlay = imageURLs.length > countFrom && countFrom == 1 ? this.renderCountOverlay(true) : this.renderOverlay();

        return <Grid>
            <Grid.Row className="no-vertical-padding">
                <Grid.Column className={`border height-two background pointer`} onClick={this.openModal.bind(this, 0)} style={{ background: `url(${imageURLs[0]})` }}>
                    {overlay}
                </Grid.Column>
            </Grid.Row>
        </Grid>;
    }

    renderTwo() {
        const { countFrom } = this.state;
        const { images } = this.props;
        const imageURLs = images.map(image => noAuthStatic(`/images/others/${image.id}?size=origin`));
        const overlay = imageURLs.length > countFrom && [2, 3].includes(+countFrom) ? this.renderCountOverlay(true) : this.renderOverlay();
        const conditionalRender = [3, 4].includes(imageURLs.length) || imageURLs.length > +countFrom && [3, 4].includes(+countFrom);

        return <Grid>
            <Grid.Row columns={2} className="no-vertical-padding">
                <Grid.Column className="border height-two background pointer" onClick={this.openModal.bind(this, conditionalRender ? 1 : 0)} style={{ background: `url(${conditionalRender ? imageURLs[1] : imageURLs[0]})` }}>
                    {this.renderOverlay()}
                </Grid.Column>
                <Grid.Column className="border height-two background pointer" onClick={this.openModal.bind(this, conditionalRender ? 2 : 1)} style={{ background: `url(${conditionalRender ? imageURLs[2] : imageURLs[1]})` }}>
                    {overlay}
                </Grid.Column>
            </Grid.Row>
        </Grid>;
    }

    renderThree() {
        const { countFrom } = this.state;
        const { images } = this.props;
        const imageURLs = images.map(image => noAuthStatic(`/images/others/${image.id}?size=origin`));
        const overlay = !countFrom || countFrom > 5 || imageURLs.length > countFrom && [4, 5].includes(+countFrom) ? this.renderCountOverlay(true) : this.renderOverlay(conditionalRender ? 3 : 4);
        const conditionalRender = imageURLs.length == 4 || imageURLs.length > +countFrom && +countFrom == 4;

        return <Grid>
            <Grid.Row columns={3} className="no-vertical-padding">
                <Grid.Column className="border height-three background pointer" onClick={this.openModal.bind(this, conditionalRender ? 1 : 2)} style={{ background: `url(${conditionalRender ? imageURLs[1] : imageURLs[2]})` }}>
                    {this.renderOverlay(conditionalRender ? 1 : 2)}
                </Grid.Column>
                <Grid.Column className="border height-three background pointer" onClick={this.openModal.bind(this, conditionalRender ? 2 : 3)} style={{ background: `url(${conditionalRender ? imageURLs[2] : imageURLs[3]})` }}>
                    {this.renderOverlay(conditionalRender ? 2 : 3)}
                </Grid.Column>
                <Grid.Column className="border height-three background pointer" onClick={this.openModal.bind(this, conditionalRender ? 3 : 4)} style={{ background: `url(${conditionalRender ? imageURLs[3] : imageURLs[4]})` }}>
                    {overlay}
                </Grid.Column>
            </Grid.Row>
        </Grid>;
    }

    // Return overlay of the image, can combine with this.props.renderOverlay to make overlay dynamically
    renderOverlay(id) {
        return null
    }

    renderCountOverlay(more) {
        const { images } = this.props;
        const { countFrom } = this.state;
        const extra = images.length - (countFrom && countFrom > 5 ? 5 : countFrom);

        return [more && <div key="count" className="cover"></div>, more && <div key="count-sub" className="cover-text" style={{ fontSize: '200%' }}><p>+{extra}</p></div>]
    }

    render() {
        const { countFrom } = this.state;
        const { images } = this.props;
        const imagesToShow = [...images];

        if (countFrom && images.length > countFrom) {
            imagesToShow.length = countFrom;
        }

        return (
            <div className="grid-container">
                {[1, 3, 4].includes(imagesToShow.length) && this.renderOne()}
                {imagesToShow.length >= 2 && imagesToShow.length != 4 && this.renderTwo()}
                {imagesToShow.length >= 4 && this.renderThree()}
            </div>
        )
    }

}

GridImageContainer.propTypes = {
    images: PropTypes.array.isRequired,
    hideOverlay: PropTypes.bool,
    renderOverlay: PropTypes.func,
    overlayBackgroundColor: PropTypes.string,
    onClickEach: PropTypes.func,
    countFrom: PropTypes.number,
};

const mapDispatchToProps = (dispatch) => ({
    openThreaterModal: (post) => dispatch(threaterActions.openModal(post))
});

export default connect(null, mapDispatchToProps)(GridImageContainer);