import { Typography, Box } from "@mui/material";
import { PROFILE } from '../Constants'

const ProfilePage = () => {
	return(
		<>
			<Box mt={15} display='flex' justifyContent='center'>
				<Typography variant='h3'>
					{PROFILE.welcome}
				</Typography>
			</Box>

			<Box display='flex' justifyContent='center'>
				<Typography variant='h5'>
					{PROFILE.subtitle}
				</Typography>
			</Box>
		</>
	);
}



export default ProfilePage;