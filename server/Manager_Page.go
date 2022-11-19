package server

type PageManager struct {
	loadedPages map[string]Page
}

func (pm *PageManager) MoveGuestToPage(user UserControllerInterface, pagePath string, server *Server) *Page {
	// see what page the guest is currently on (user.onPage)
	// tell that page to remove the guest from its map

	// does the cached page already exist?
	// - if so add the guest to the guest map for that page
	// - if not, create the page by querying the database (server.db.query)
	// -- add the page to the loadedPages map
	// -- add the guest to the guest map for that page
	return nil
}

func (pm *PageManager) MoveMemberToPage(user UserControllerInterface, pagePath string, server *Server) {

}

func (pm *PageManager) UnloadEmptyPages(server *Server) {
	// iterate through every page
	// - if the guests map and the members map on a page is empty...
	// -- remove that page, or at least remove the reference to it from the loadedPages map so that the garbage collector will remove it.
}
