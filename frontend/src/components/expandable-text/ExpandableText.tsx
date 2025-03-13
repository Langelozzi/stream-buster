import { Button } from "@mui/material";
import { useState } from "react";

interface ExpandableTextProps {
	text: String,
	maxLength: number,
}
// Helper component for expandable text
export const ExpandableText: React.FC<ExpandableTextProps> = ({ text, maxLength, }) => {
	const [expanded, setExpanded] = useState(false);

	return (
		<>
			{expanded ? text : `${text.slice(0, maxLength)}...`}
			< Button
				onClick={() => setExpanded(!expanded)}
				sx={{ ml: 1, minWidth: 'auto', p: '2px 5px', fontSize: '0.75rem' }}
			>
				{expanded ? 'Show Less' : 'Read More'}
			</Button >
		</>
	);
};

