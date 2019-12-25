package main

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"os"
)

type roleInfo struct {
	Role      string
	User      string
	Namespace string
}
var roles []roleInfo

func UserToRoles(clientset *kubernetes.Clientset, user string, outputKind string){
	addClusterRoleBindings(clientset, user, "")
	addRoleBindings(clientset,user,"")
	outputData(outputKind,"User")
}

func RolesToUser(clientset *kubernetes.Clientset, role string, outputKind string){
	addClusterRoleBindings(clientset, "", role)
	addRoleBindings(clientset, "", role)
	outputData(outputKind,"Role")
}

func addClusterRoleBindings(clientset *kubernetes.Clientset, user string, role string){
	var clusterRoleBindings, err = clientset.RbacV1().ClusterRoleBindings().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	} else {
		//fmt.Print(clusterRoleBindings)
		for _, rolebinding := range clusterRoleBindings.Items {
			for _, subjects := range rolebinding.Subjects {
				if subjects.Kind == "User" {
					if user == "" || subjects.Name == user {
						if role == "" || rolebinding.RoleRef.Name == role {
							roles = append(roles, roleInfo{Role: rolebinding.RoleRef.Name, Namespace: rolebinding.ObjectMeta.Namespace, User: subjects.Name})
						}
					}
				}
			}
		}
	}
}
func addRoleBindings(clientset *kubernetes.Clientset, user string, role string){
	var roleBindings, err = clientset.RbacV1().RoleBindings("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	} else {
		//fmt.Print(clusterRoleBindings)
		for _, rolebinding := range roleBindings.Items {
			for _, subjects := range rolebinding.Subjects {
				if subjects.Kind == "User" {
					if user == "" || subjects.Name == user {
						if role == "" || rolebinding.RoleRef.Name == role {
							roles = append(roles, roleInfo{Role: rolebinding.RoleRef.Name, Namespace: rolebinding.ObjectMeta.Namespace, User: subjects.Name})
						}
					}
				}
			}
		}
	}
}
func outputData(outputType string, key string){
	if outputType == "table" {
		tableOutput(key)
	}else{
		jsonOutput()
	}
}

func tableOutput(key string) {
	data := make([][]string, 0)
	table := tablewriter.NewWriter(os.Stdout)
	if key == "User"{
		table.SetHeader([]string{"Username", "Role", "Namespace"})
		for _, value := range roles{
			data = append(data, []string{value.User, value.Role, value.Namespace})
		}
	}else if key == "Role" {
		table.SetHeader([]string{"Role", "Username", "Namespace"})
		for _, value := range roles{
			data = append(data, []string{value.Role, value.User, value.Namespace})
		}
	}else {
		log.Fatal("Invalid output key")
	}

	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}

func jsonOutput() {
	outputJson, _ := json.MarshalIndent(roles, "", "  ")
	os.Stdout.Write(outputJson)
}