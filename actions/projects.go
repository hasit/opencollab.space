package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/hasit/opencollab.space/models"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Project)
// DB Table: Plural (projects)
// Resource: Plural (Projects)
// Path: Plural (/projects)
// View Template Folder: Plural (/templates/projects/)

// ProjectsResource is the resource for the project model
type ProjectsResource struct {
	buffalo.Resource
}

// List gets all Projects. This function is mapped to the path
// GET /projects
func (v ProjectsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	projects := &models.Projects{}
	// You can order your list here. Just change
	err := tx.All(projects)
	// to:
	// err := tx.Order("create_at desc").All(projects)
	if err != nil {
		return errors.WithStack(err)
	}
	// Make Projects available inside the html template
	c.Set("projects", projects)
	return c.Render(200, r.HTML("projects/index.html"))
}

// Show gets the data for one Project. This function is mapped to
// the path GET /projects/{project_id}
func (v ProjectsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Project
	project := &models.Project{}
	// To find the Project the parameter project_id is used.
	err := tx.Find(project, c.Param("project_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Make project available inside the html template
	c.Set("project", project)
	return c.Render(200, r.HTML("projects/show.html"))
}

// New renders the formular for creating a new Project.
// This function is mapped to the path GET /projects/new
func (v ProjectsResource) New(c buffalo.Context) error {
	// Make project available inside the html template
	c.Set("project", &models.Project{})
	return c.Render(200, r.HTML("projects/new.html"))
}

// Create adds a Project to the DB. This function is mapped to the
// path POST /projects
func (v ProjectsResource) Create(c buffalo.Context) error {
	// Allocate an empty Project
	project := &models.Project{}
	// Bind project to the html form elements
	err := c.Bind(project)
	if err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(project)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		// Make project available inside the html template
		c.Set("project", project)
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("projects/new.html"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "Project was created successfully")
	// and redirect to the projects index page
	return c.Redirect(302, "/projects/%s", project.ID)
}

// Edit renders a edit formular for a project. This function is
// mapped to the path GET /projects/{project_id}/edit
func (v ProjectsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Project
	project := &models.Project{}
	err := tx.Find(project, c.Param("project_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Make project available inside the html template
	c.Set("project", project)
	return c.Render(200, r.HTML("projects/edit.html"))
}

// Update changes a project in the DB. This function is mapped to
// the path PUT /projects/{project_id}
func (v ProjectsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Project
	project := &models.Project{}
	err := tx.Find(project, c.Param("project_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Bind Project to the html form elements
	err = c.Bind(project)
	if err != nil {
		return errors.WithStack(err)
	}
	verrs, err := tx.ValidateAndUpdate(project)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		// Make project available inside the html template
		c.Set("project", project)
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("projects/edit.html"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "Project was updated successfully")
	// and redirect to the projects index page
	return c.Redirect(302, "/projects/%s", project.ID)
}

// Destroy deletes a project from the DB. This function is mapped
// to the path DELETE /projects/{project_id}
func (v ProjectsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Project
	project := &models.Project{}
	// To find the Project the parameter project_id is used.
	err := tx.Find(project, c.Param("project_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	err = tx.Destroy(project)
	if err != nil {
		return errors.WithStack(err)
	}
	// If there are no errors set a flash message
	c.Flash().Add("success", "Project was destroyed successfully")
	// Redirect to the projects index page
	return c.Redirect(302, "/projects")
}
