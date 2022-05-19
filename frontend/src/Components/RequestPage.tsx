import { Typography, Box } from "@mui/material";
import { REQUEST } from '../Constants'

const RequestPage = () => {
	return(
		<>
			<Box mt={15} display='flex' justifyContent='center'>
				<Typography variant='h3'>
					{REQUEST.welcome}
				</Typography>
			</Box>

			<Box display='flex' justifyContent='center'>
				<Typography variant='h5'>
					{REQUEST.subtitle}
				</Typography>
			</Box>
		</>
	);
}

export default RequestPage;