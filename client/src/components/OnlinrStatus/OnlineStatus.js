import React, { Component, Fragment } from 'react';
import { Icon } from 'semantic-ui-react';

import { getTimeDiffFromNow } from '../../utils/time';
import Socket from '../../Socket';

class OnlineStatus extends Component {
    constructor(props) {
        super(props);

        this.state = {
            lastLogin: props.lastLogin,
            lastLogout: props.lastLogout,
            displayText: this.getDisplayText(props.lastLogin, props.lastLogout)
        }
    }

    componentDidMount() {
        this.setDisplayTime();
        this.interval = setInterval(this.setDisplayTime, 5000);

        Socket.on("USER.SERVER.GO_OFFLINE", this.setUserOffline);
        Socket.on("USER.SERVER.GO_ONLINE", this.setUserOnline);
    }

    componentWillUnmount() {
        Socket.removeListener("USER.SERVER.GO_OFFLINE", this.setUserOffline);
        Socket.removeListener("USER.SERVER.GO_ONLINE", this.setUserOnline);
    }

    componentWillReceiveProps(nextProps) {
        if (this.props.lastLogin !== nextProps.lastLogin) {
            this.setState({ lastLogin: nextProps.lastLogin });
        }

        if (this.props.lastLogout !== nextProps.lastLogout) {
            this.setState({ lastLogout: nextProps.lastLogout });
        }
    }

    setUserOffline = (message) => {
        const { actorID } = message;
        const { partnerID } = this.props;

        if (partnerID === actorID) {
            this.setState({ lastLogout: new Date() })
        }
    }

    setUserOnline = (message) => {
        const { actorID } = message;
        const { partnerID } = this.props;

        if (partnerID === actorID) {
            this.setState({ lastLogin: new Date() })
        }
    }

    componentDidUpdate(_, prevState) {
        if (this.state.lastLogin !== prevState.lastLogin || this.state.lastLogout !== prevState.lastLogout) {
            clearInterval(this.interval);
            this.setDisplayTime();
            this.interval = setInterval(this.setDisplayTime, 5000);
        }
    }

    setDisplayTime = () => {
        const { lastLogin, lastLogout } = this.state;
        this.setState({ displayText: this.getDisplayText(lastLogin, lastLogout) });
    }

    componentWillUnmount = () => {
        clearInterval(this.interval);
    }

    getDisplayText(lastLogin, lastLogout) {
        if (lastLogout == null || new Date(lastLogin) > new Date(lastLogout)) {
            return "Online";
        }

        return getTimeDiffFromNow(lastLogout).toString();
    }

    render() {
        const { displayText } = this.state;

        return (
            <Fragment>
                <Icon name="circle" size="mini" color={displayText === "Online" ? "green" : "grey"} />
                <span>{displayText}</span>
            </Fragment>
        );
    }
}

export default OnlineStatus;