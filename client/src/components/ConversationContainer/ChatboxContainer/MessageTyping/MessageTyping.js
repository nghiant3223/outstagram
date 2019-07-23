import React from 'react'

export default function MessageTyping() {
    return (
        <div className={`lhs-message-container`}>
            <div className="lhs-message-avatar">
                <img src={user.avatarUrl} alt="Avatar" />
            </div>
            <div className="lhs-message-container__messages">
                <div className={`lhs-message-item`}>
                    <div className={`lhs-message-item__content`}>
                        <div className={`message-typing message-typing--lhs`}>
                            <div></div>
                            <div></div>
                            <div></div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}
