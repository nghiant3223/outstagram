export function groupChatName(members) {
    return members.map((member) => member.fullname).join(", ");
}