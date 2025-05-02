
type TextTruncateProps = {
    text: string;
    maxLength?: number;
  };

const TextTruncate = ({ text, maxLength = 60 } : TextTruncateProps) => {
    if (text.length <= maxLength) return text;
    return `${text.substring(0, maxLength)}...`;
  };

export default TextTruncate