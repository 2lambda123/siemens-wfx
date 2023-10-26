// SPDX-FileCopyrightText: The entgo authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// HistoryColumns holds the columns for the "history" table.
	HistoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "mtime", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "TIMESTAMP(6)"}},
		{Name: "status", Type: field.TypeJSON, Nullable: true},
		{Name: "definition", Type: field.TypeJSON, Nullable: true},
		{Name: "job_history", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// HistoryTable holds the schema information for the "history" table.
	HistoryTable = &schema.Table{
		Name:       "history",
		Columns:    HistoryColumns,
		PrimaryKey: []*schema.Column{HistoryColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "history_job_history",
				Columns:    []*schema.Column{HistoryColumns[4]},
				RefColumns: []*schema.Column{JobColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// JobColumns holds the columns for the "job" table.
	JobColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "stime", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "TIMESTAMP(6)"}},
		{Name: "mtime", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "TIMESTAMP(6)"}},
		{Name: "client_id", Type: field.TypeString},
		{Name: "definition", Type: field.TypeJSON, Nullable: true},
		{Name: "status", Type: field.TypeJSON},
		{Name: "group", Type: field.TypeString, Nullable: true},
		{Name: "workflow_jobs", Type: field.TypeInt, Nullable: true},
	}
	// JobTable holds the schema information for the "job" table.
	JobTable = &schema.Table{
		Name:       "job",
		Columns:    JobColumns,
		PrimaryKey: []*schema.Column{JobColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "job_workflow_jobs",
				Columns:    []*schema.Column{JobColumns[7]},
				RefColumns: []*schema.Column{WorkflowColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "job_stime",
				Unique:  false,
				Columns: []*schema.Column{JobColumns[1]},
			},
			{
				Name:    "job_client_id",
				Unique:  false,
				Columns: []*schema.Column{JobColumns[3]},
			},
			{
				Name:    "job_group",
				Unique:  false,
				Columns: []*schema.Column{JobColumns[6]},
			},
		},
	}
	// TagColumns holds the columns for the "tag" table.
	TagColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// TagTable holds the schema information for the "tag" table.
	TagTable = &schema.Table{
		Name:       "tag",
		Columns:    TagColumns,
		PrimaryKey: []*schema.Column{TagColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "tag_name",
				Unique:  true,
				Columns: []*schema.Column{TagColumns[1]},
			},
		},
	}
	// WorkflowColumns holds the columns for the "workflow" table.
	WorkflowColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 1024},
		{Name: "states", Type: field.TypeJSON},
		{Name: "transitions", Type: field.TypeJSON},
		{Name: "groups", Type: field.TypeJSON},
	}
	// WorkflowTable holds the schema information for the "workflow" table.
	WorkflowTable = &schema.Table{
		Name:       "workflow",
		Columns:    WorkflowColumns,
		PrimaryKey: []*schema.Column{WorkflowColumns[0]},
	}
	// TagJobsColumns holds the columns for the "tag_jobs" table.
	TagJobsColumns = []*schema.Column{
		{Name: "tag_id", Type: field.TypeInt},
		{Name: "job_id", Type: field.TypeString, Size: 36},
	}
	// TagJobsTable holds the schema information for the "tag_jobs" table.
	TagJobsTable = &schema.Table{
		Name:       "tag_jobs",
		Columns:    TagJobsColumns,
		PrimaryKey: []*schema.Column{TagJobsColumns[0], TagJobsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tag_jobs_tag_id",
				Columns:    []*schema.Column{TagJobsColumns[0]},
				RefColumns: []*schema.Column{TagColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tag_jobs_job_id",
				Columns:    []*schema.Column{TagJobsColumns[1]},
				RefColumns: []*schema.Column{JobColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		HistoryTable,
		JobTable,
		TagTable,
		WorkflowTable,
		TagJobsTable,
	}
)

func init() {
	HistoryTable.ForeignKeys[0].RefTable = JobTable
	HistoryTable.Annotation = &entsql.Annotation{
		Table: "history",
	}
	JobTable.ForeignKeys[0].RefTable = WorkflowTable
	JobTable.Annotation = &entsql.Annotation{
		Table: "job",
	}
	TagTable.Annotation = &entsql.Annotation{
		Table: "tag",
	}
	WorkflowTable.Annotation = &entsql.Annotation{
		Table: "workflow",
	}
	TagJobsTable.ForeignKeys[0].RefTable = TagTable
	TagJobsTable.ForeignKeys[1].RefTable = JobTable
}
