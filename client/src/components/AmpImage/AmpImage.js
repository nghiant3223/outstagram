import React from 'react';
import PropTypes from 'prop-types'

import "./AmpImage.css";

function AmpImage(props) {
    const { fitType, src, container } = props;

    const style = {};
    if (typeof container === "object") {
        style.width = container.width;
        style.height = container.height;
    }

    return (
        <div className={container === "auto" ? "AutoResizeContainer" : "FixedContainer"} style={{ ...style }}>
            <amp-img class={fitType === "contain" ? "Contain" : "Cover"} layout="fill" src={src}></amp-img>
        </div>
    );
}

AmpImage.propTypes = {
    fitType: PropTypes.oneOf(["contain", "cover"]).isRequired,
    src: PropTypes.string.isRequired,
    container: PropTypes.oneOfType([PropTypes.oneOf(["auto"]), PropTypes.shape({ width: PropTypes.number.isRequired, height: PropTypes.number.isRequired })])
}

AmpImage.defaultProps = {
    container: "auto"
}

export default AmpImage;