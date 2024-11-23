
type ComparableObject = {
	UpdatedAt: string; // ISO 8601 date format: "YYYY-MM-DDTHH:mm:ss.sss±HH:mm"
};
export const getFormattedDate = (): string => {
	const date = new Date();

	// Get components of the date
	const year = date.getFullYear();
	const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are 0-based
	const day = String(date.getDate()).padStart(2, '0');
	const hours = String(date.getHours()).padStart(2, '0');
	const minutes = String(date.getMinutes()).padStart(2, '0');
	const seconds = String(date.getSeconds()).padStart(2, '0');
	const milliseconds = String(date.getMilliseconds()).padStart(3, '0');

	// Get timezone offset in minutes
	const offset = -date.getTimezoneOffset();
	const absOffset = Math.abs(offset);
	const offsetHours = String(Math.floor(absOffset / 60)).padStart(2, '0');
	const offsetMinutes = String(absOffset % 60).padStart(2, '0');
	const timezone = `${offset >= 0 ? '+' : '-'}${offsetHours}:${offsetMinutes}`;

	return `${year}-${month}-${day}T${hours}:${minutes}:${seconds}.${milliseconds}${timezone}`;
};

export const compareByUpdatedAt = (a: ComparableObject, b: ComparableObject): number => {
	const dateA = new Date(a.UpdatedAt);
	const dateB = new Date(b.UpdatedAt);

	if (dateA < dateB) return 1; // a is earlier
	if (dateA > dateB) return -1;  // a is later
	return 0; // dates are equal
};
