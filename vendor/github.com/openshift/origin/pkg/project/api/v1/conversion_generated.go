// +build !ignore_autogenerated

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	project_api "github.com/openshift/origin/pkg/project/api"
	api "k8s.io/kubernetes/pkg/api"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	reflect "reflect"
)

func init() {
	if err := api.Scheme.AddGeneratedConversionFuncs(
		Convert_v1_Project_To_api_Project,
		Convert_api_Project_To_v1_Project,
		Convert_v1_ProjectList_To_api_ProjectList,
		Convert_api_ProjectList_To_v1_ProjectList,
		Convert_v1_ProjectRequest_To_api_ProjectRequest,
		Convert_api_ProjectRequest_To_v1_ProjectRequest,
		Convert_v1_ProjectSpec_To_api_ProjectSpec,
		Convert_api_ProjectSpec_To_v1_ProjectSpec,
		Convert_v1_ProjectStatus_To_api_ProjectStatus,
		Convert_api_ProjectStatus_To_v1_ProjectStatus,
	); err != nil {
		// if one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}
}

func autoConvert_v1_Project_To_api_Project(in *Project, out *project_api.Project, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*Project))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v1_ProjectSpec_To_api_ProjectSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ProjectStatus_To_api_ProjectStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_Project_To_api_Project(in *Project, out *project_api.Project, s conversion.Scope) error {
	return autoConvert_v1_Project_To_api_Project(in, out, s)
}

func autoConvert_api_Project_To_v1_Project(in *project_api.Project, out *Project, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*project_api.Project))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_api_ProjectSpec_To_v1_ProjectSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_api_ProjectStatus_To_v1_ProjectStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_Project_To_v1_Project(in *project_api.Project, out *Project, s conversion.Scope) error {
	return autoConvert_api_Project_To_v1_Project(in, out, s)
}

func autoConvert_v1_ProjectList_To_api_ProjectList(in *ProjectList, out *project_api.ProjectList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*ProjectList))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]project_api.Project, len(*in))
		for i := range *in {
			if err := Convert_v1_Project_To_api_Project(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_ProjectList_To_api_ProjectList(in *ProjectList, out *project_api.ProjectList, s conversion.Scope) error {
	return autoConvert_v1_ProjectList_To_api_ProjectList(in, out, s)
}

func autoConvert_api_ProjectList_To_v1_ProjectList(in *project_api.ProjectList, out *ProjectList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*project_api.ProjectList))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Project, len(*in))
		for i := range *in {
			if err := Convert_api_Project_To_v1_Project(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_ProjectList_To_v1_ProjectList(in *project_api.ProjectList, out *ProjectList, s conversion.Scope) error {
	return autoConvert_api_ProjectList_To_v1_ProjectList(in, out, s)
}

func autoConvert_v1_ProjectRequest_To_api_ProjectRequest(in *ProjectRequest, out *project_api.ProjectRequest, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*ProjectRequest))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	return nil
}

func Convert_v1_ProjectRequest_To_api_ProjectRequest(in *ProjectRequest, out *project_api.ProjectRequest, s conversion.Scope) error {
	return autoConvert_v1_ProjectRequest_To_api_ProjectRequest(in, out, s)
}

func autoConvert_api_ProjectRequest_To_v1_ProjectRequest(in *project_api.ProjectRequest, out *ProjectRequest, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*project_api.ProjectRequest))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	return nil
}

func Convert_api_ProjectRequest_To_v1_ProjectRequest(in *project_api.ProjectRequest, out *ProjectRequest, s conversion.Scope) error {
	return autoConvert_api_ProjectRequest_To_v1_ProjectRequest(in, out, s)
}

func autoConvert_v1_ProjectSpec_To_api_ProjectSpec(in *ProjectSpec, out *project_api.ProjectSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*ProjectSpec))(in)
	}
	if in.Finalizers != nil {
		in, out := &in.Finalizers, &out.Finalizers
		*out = make([]api.FinalizerName, len(*in))
		for i := range *in {
			(*out)[i] = api.FinalizerName((*in)[i])
		}
	} else {
		out.Finalizers = nil
	}
	return nil
}

func Convert_v1_ProjectSpec_To_api_ProjectSpec(in *ProjectSpec, out *project_api.ProjectSpec, s conversion.Scope) error {
	return autoConvert_v1_ProjectSpec_To_api_ProjectSpec(in, out, s)
}

func autoConvert_api_ProjectSpec_To_v1_ProjectSpec(in *project_api.ProjectSpec, out *ProjectSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*project_api.ProjectSpec))(in)
	}
	if in.Finalizers != nil {
		in, out := &in.Finalizers, &out.Finalizers
		*out = make([]api_v1.FinalizerName, len(*in))
		for i := range *in {
			(*out)[i] = api_v1.FinalizerName((*in)[i])
		}
	} else {
		out.Finalizers = nil
	}
	return nil
}

func Convert_api_ProjectSpec_To_v1_ProjectSpec(in *project_api.ProjectSpec, out *ProjectSpec, s conversion.Scope) error {
	return autoConvert_api_ProjectSpec_To_v1_ProjectSpec(in, out, s)
}

func autoConvert_v1_ProjectStatus_To_api_ProjectStatus(in *ProjectStatus, out *project_api.ProjectStatus, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*ProjectStatus))(in)
	}
	out.Phase = api.NamespacePhase(in.Phase)
	return nil
}

func Convert_v1_ProjectStatus_To_api_ProjectStatus(in *ProjectStatus, out *project_api.ProjectStatus, s conversion.Scope) error {
	return autoConvert_v1_ProjectStatus_To_api_ProjectStatus(in, out, s)
}

func autoConvert_api_ProjectStatus_To_v1_ProjectStatus(in *project_api.ProjectStatus, out *ProjectStatus, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*project_api.ProjectStatus))(in)
	}
	out.Phase = api_v1.NamespacePhase(in.Phase)
	return nil
}

func Convert_api_ProjectStatus_To_v1_ProjectStatus(in *project_api.ProjectStatus, out *ProjectStatus, s conversion.Scope) error {
	return autoConvert_api_ProjectStatus_To_v1_ProjectStatus(in, out, s)
}
