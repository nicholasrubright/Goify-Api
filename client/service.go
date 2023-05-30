package client

import (
	"github.com/goify-api/client/models"
	spotify "github.com/goify-api/spotify"
	"github.com/goify-api/utils"
)

// Interacts with the spotify service and the frontend client

func GetAuthorizationUrl(clientId string, redirectUri string) (*models.ClientAuthorizationUrlResponse, *models.ClientErrorResponse) {

	spotifyAuthorizationResponse, err := spotify.GetAuthorizationUrl(clientId, redirectUri)

	if err != nil {
		utils.SendEndpointError("GetAuthorizationUrl", err)
		return nil, &models.ClientErrorResponse{
			Error: "GetAuthorizationUrl",
			Message: "There was a problem creating the authorization url",
		}
	}

	return &models.ClientAuthorizationUrlResponse{
		Url: spotifyAuthorizationResponse.Url,
	}, nil

} 

func GetAccessToken(clientId string, clientSecret string, code string, redirectUri string) (*models.ClientTokenResponse, *models.ClientErrorResponse) {

	spotifyAccessTokenResponse, spotifyAuthorizationErrorResponse, err := spotify.GetAccessToken(clientId, clientSecret, code, redirectUri)


	if err != nil || spotifyAuthorizationErrorResponse != nil {
		utils.SendEndpointError("GetAccessToken", err)
		return nil, &models.ClientErrorResponse{
			Error: "GetAccessToken",
			Message: "There was a problem getting the access token",
		}
	}

	return &models.ClientTokenResponse{
		Token: spotifyAccessTokenResponse.Token,
		ExpiresIn: spotifyAccessTokenResponse.Expires,
	}, nil

}

func GetCurrentUserProfile(token string) (*models.ClientUserProfileResponse, *models.ClientErrorResponse) {
	spotifyCurrentUserProfileResponse, spotifyErrorResponse, err := spotify.GetCurrentUserProfile(token)

	if err != nil || spotifyErrorResponse != nil {
		utils.SendEndpointError("GetCurrentUserProfile", err)
		return nil, &models.ClientErrorResponse{
			Error: "GetCurrentUserProfile",
			Message: "There was a problem getting the current user profile",
		}
	}

	var clientImages []models.ClientImageResponse

	clientImages = make([]models.ClientImageResponse, 0)


	for _, spotifyImageResponse := range spotifyCurrentUserProfileResponse.Images {
		newClientImage := models.ClientImageResponse{
			URL: spotifyImageResponse.URL,
			Height: spotifyImageResponse.Height,
			Width: spotifyImageResponse.Width,
		}
		clientImages = append(clientImages, newClientImage)
	}

	clientCurrentUserProfileResponse := models.ClientUserProfileResponse{
		ID: spotifyCurrentUserProfileResponse.ID,
		Name: spotifyCurrentUserProfileResponse.DisplayName,
		Images: clientImages,
	}
	return &clientCurrentUserProfileResponse, nil
}

func GetCurrentUserPlaylists(token string) (*models.ClientUserPlaylistsResponse, *models.ClientErrorResponse) {
	spotifyCurrentUserPlaylistsResponse, spotifyErrorResponse, err := spotify.GetCurrentUserPlaylists(token)

	if err != nil || spotifyErrorResponse != nil {
		utils.SendEndpointError("GetCurrentUserPlaylists", err)
		return nil, &models.ClientErrorResponse{
			Error: "GetCurrentUserPlaylists",
			Message: "There was a problem getting the current user playlists",
		}
	}

	var clientPlaylists []models.ClientPlaylistResponse

	clientPlaylists = make([]models.ClientPlaylistResponse, 0)

	for _, spotifyPlaylistsResponse := range spotifyCurrentUserPlaylistsResponse.Items {
		
		var clientPlaylistImage []models.ClientImageResponse
		clientPlaylistImage = make([]models.ClientImageResponse, 0)


		for _, spotifyPlaylistImageResponse := range spotifyPlaylistsResponse.Images {
			newClientImage := models.ClientImageResponse {
				URL: spotifyPlaylistImageResponse.URL,
				Height: spotifyPlaylistImageResponse.Height,
				Width: spotifyPlaylistImageResponse.Width,
			}
			clientPlaylistImage = append(clientPlaylistImage, newClientImage)
		}		
		
		newClientPlaylist := models.ClientPlaylistResponse{
			ID: spotifyPlaylistsResponse.ID,
			Images: clientPlaylistImage,
			Name: spotifyPlaylistsResponse.Name,
		}

		clientPlaylists = append(clientPlaylists, newClientPlaylist)

	}

	return &models.ClientUserPlaylistsResponse {
		Limit: spotifyCurrentUserPlaylistsResponse.Limit,
		Next: spotifyCurrentUserPlaylistsResponse.Next,
		Offset: spotifyCurrentUserPlaylistsResponse.Offset,
		Previous: spotifyCurrentUserPlaylistsResponse.Previous,
		Total: spotifyCurrentUserPlaylistsResponse.Total,
		Playlists: clientPlaylists,
	}, nil
}

/*	Combine Playlists
*	1. Create blank Playlist with name and description
*	2. Get all tracks from the selected playlists
*	3. Add each track to the new playlist
*	4. Return new playlist name and description
*/
func CombinePlaylists(playlist_name string, playlist_desc string, playlist_ids []string, token string) {

}