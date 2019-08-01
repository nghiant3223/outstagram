const SECOND = 1000;
const MINUTE = SECOND * 60;
const HOUR = MINUTE * 60;
const DAY = HOUR * 24;

export function getTimeDiffFromNow(time) {
    const then = new Date(time);
    const now = new Date();
    const diff = now.getTime() - then.getTime();

    if (diff < HOUR) {
        return (parseInt(diff / MINUTE, 10) + 1).toString() + 'min';
    }
    else if (diff < 24 * HOUR) {
        return (parseInt(diff / HOUR, 10)).toString() + 'hr';
    }
    else if (diff < 7 * DAY) {
        return (parseInt(diff / DAY, 10)).toString() + 'd';
    } else {
        return time;
    }
}

export function toHHMM(time) {
    const h = time.getHours();
    const m = time.getMinutes();
    return (h < 10 ? ('0' + h.toString()) : h.toString()) + ':' + (m < 10 ? ('0' + m.toString()) : m.toString());
}

